import { BrowserRouter, Routes, Route } from "react-router-dom";
import RootLayout from "@/layouts/RootLayout";
import AppLayout from "@/layouts/AppLayout";
import AuthLayout from "@/layouts/AuthLayout";

import OverviewView from "@/views/workspace/overview/OverviewView";
import OnboardingView from "@/views/onboarding/OnboardingView";
import LoginView from "@/views/auth/LoginView";
import NotFoundView from "@/views/NotFoundView";
import NodeView from "./views/workspace/nodes/NodeView";

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
            <Route path="nodes" element={<NodeView />} />


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
