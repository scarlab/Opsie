import { BrowserRouter, Routes, Route } from "react-router-dom";
import RootLayout from "@/layouts/RootLayout";
import AppLayout from "@/layouts/AppLayout";
import AuthLayout from "@/layouts/AuthLayout";

import OverviewView from "@/views/workspace/overview/OverviewView";
import OnboardingView from "@/views/onboarding/OnboardingView";
import LoginView from "@/views/auth/LoginView";
import NotFoundView from "@/views/NotFoundView";
import NodesView from "./views/workspace/nodes/NodesView";
import AppsView from "./views/workspace/apps/AppsView";
import ResourcesView from "./views/workspace/resources/ResourcesView";
import ProjectsView from "./views/workspace/projects/ProjectsView";
import UsersView from "./views/workspace/users/UsersView";
import SettingsLayout from "./layouts/SettingsLayout";
import OrganizationSettingsView from "./views/settings/organization/OrganizationSettingsView";
import AccountSettingsView from "./views/settings/account/AccountSettingsView";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Root shell (common wrappers, providers, etc.) */}
        <Route element={<RootLayout />}>

          {/* =====================
              Onboarding routes
          ====================== */}
          <Route path="onboarding" element={<OnboardingView />} />

          {/* =====================
              Auth routes
          ====================== */}
          <Route path="auth/*" element={<AuthLayout />}>
            <Route path="login" element={<LoginView />} />


            {/* fallback for /auth or /auth/anything */}
            <Route path="*" element={<NotFoundView />} />
          </Route>




          {/* =====================
              App routes (protected)
          ====================== */}
          <Route path="/*" element={<AppLayout />}>
            <Route index element={<OverviewView />} />
            <Route path="nodes" element={<NodesView />} />
            <Route path="users" element={<UsersView />} />
            <Route path="projects" element={<ProjectsView />} />
            <Route path="resources" element={<ResourcesView />} />
            <Route path="apps" element={<AppsView />} />


            {/* Add more nested routes later */}
            <Route path="*" element={<div className=""><NotFoundView /></div>} />
          </Route>

          {/* =====================
              App routes (protected)
          ====================== */}
          <Route path="settings/*" element={<SettingsLayout />}>
            <Route index element={<OrganizationSettingsView />} />
            <Route path="account" element={<AccountSettingsView />} />


            {/* Add more nested routes later */}
            <Route path="*" element={<div className=""><NotFoundView /></div>} />
          </Route>




          {/* =====================
              Global 404 (outside layouts)
          ====================== */}
          <Route path="*" element={<NotFoundView />} />

        </Route>
      </Routes>
    </BrowserRouter>
  );
}
