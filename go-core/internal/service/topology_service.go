package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Node struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Type   string `json:"type"`
}

type TopologyService struct {
	driver neo4j.DriverWithContext
}

func NewTopologyService(driver neo4j.DriverWithContext) *TopologyService {
	return &TopologyService{driver: driver}
}

func (s *TopologyService) GetTopology(ctx context.Context, ciID uuid.UUID, depth int) ([]Node, []Edge, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
		MATCH path = (start:CI {id: $id})-[*1..$depth]-(related)
		RETURN start, related, relationships(path)
	`, map[string]interface{}{"id": ciID.String(), "depth": depth})
	if err != nil {
		return nil, nil, err
	}

	nodeMap := map[string]Node{}
	edges := make([]Edge, 0)

	for result.Next(ctx) {
		record := result.Record()
		startVal, _ := record.Get("start")
		relatedVal, _ := record.Get("related")
		relsVal, _ := record.Get("relationships")

		if n, ok := startVal.(neo4j.Node); ok {
			id, _ := n.Props["id"].(string)
			name, _ := n.Props["name"].(string)
			nodeType, _ := n.Props["type"].(string)
			status, _ := n.Props["status"].(string)
			nodeMap[id] = Node{ID: id, Name: name, Type: nodeType, Status: status}
		}
		if n, ok := relatedVal.(neo4j.Node); ok {
			id, _ := n.Props["id"].(string)
			name, _ := n.Props["name"].(string)
			nodeType, _ := n.Props["type"].(string)
			status, _ := n.Props["status"].(string)
			nodeMap[id] = Node{ID: id, Name: name, Type: nodeType, Status: status}
		}

		switch rels := relsVal.(type) {
		case []interface{}:
			for _, raw := range rels {
				if rel, ok := raw.(neo4j.Relationship); ok {
					edges = append(edges, Edge{
						Source: rel.StartElementId,
						Target: rel.EndElementId,
						Type:   rel.Type,
					})
				}
			}
		}
	}
	if err := result.Err(); err != nil {
		return nil, nil, err
	}

	nodes := make([]Node, 0, len(nodeMap))
	for _, n := range nodeMap {
		nodes = append(nodes, n)
	}
	return nodes, edges, nil
}

func (s *TopologyService) ImpactAnalysis(ctx context.Context, appID uuid.UUID, depth int) ([]Node, []Edge, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
		MATCH (app:CI {id: $appId, type: "application"})
		MATCH path = (app)-[*1..$depth]-(affected)
		WHERE affected.status = 'active'
		RETURN nodes(path) AS nodes, relationships(path) AS rels
	`, map[string]interface{}{"appId": appID.String(), "depth": depth})
	if err != nil {
		return nil, nil, err
	}

	nodeMap := map[string]Node{}
	edges := make([]Edge, 0)

	for result.Next(ctx) {
		record := result.Record()
		nodesVal, _ := record.Get("nodes")
		relsVal, _ := record.Get("rels")

		if nodeSlice, ok := nodesVal.([]interface{}); ok {
			for _, raw := range nodeSlice {
				if n, ok := raw.(neo4j.Node); ok {
					id, _ := n.Props["id"].(string)
					name, _ := n.Props["name"].(string)
					nodeType, _ := n.Props["type"].(string)
					status, _ := n.Props["status"].(string)
					nodeMap[id] = Node{ID: id, Name: name, Type: nodeType, Status: status}
				}
			}
		}

		if relSlice, ok := relsVal.([]interface{}); ok {
			for _, raw := range relSlice {
				if rel, ok := raw.(neo4j.Relationship); ok {
					edges = append(edges, Edge{
						Source: rel.StartElementId,
						Target: rel.EndElementId,
						Type:   rel.Type,
					})
				}
			}
		}
	}
	if err := result.Err(); err != nil {
		return nil, nil, err
	}

	nodes := make([]Node, 0, len(nodeMap))
	for _, n := range nodeMap {
		nodes = append(nodes, n)
	}
	return nodes, edges, nil
}
