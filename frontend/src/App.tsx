import { Navigate, Route, Routes } from "react-router-dom";
import AppLayout from "./components/AppLayout";
import ResourcesPage from "./pages/ResourcesPage";
import ResourceDetailPage from "./pages/ResourceDetailPage";
import TopologyPage from "./pages/TopologyPage";
import CiTypesPage from "./pages/CiTypesPage";
import TasksPage from "./pages/TasksPage";
import UsersPage from "./pages/UsersPage";
import LoginPage from "./pages/LoginPage";

function App() {
  return (
    <Routes>
      <Route path="/login" element={<LoginPage />} />
      <Route element={<AppLayout />}>
        <Route path="/" element={<Navigate to="/resources" replace />} />
        <Route path="/resources" element={<ResourcesPage />} />
        <Route path="/resources/:id" element={<ResourceDetailPage />} />
        <Route path="/topology/:id" element={<TopologyPage />} />
        <Route path="/ci-types" element={<CiTypesPage />} />
        <Route path="/tasks" element={<TasksPage />} />
        <Route path="/users" element={<UsersPage />} />
      </Route>
    </Routes>
  );
}

export default App;
