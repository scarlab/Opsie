import { Outlet, Link } from "react-router-dom";

export default function WorkspaceLayout() {
    return (
        <div className="flex min-h-screen">
            {/* Sidebar */}
            <aside className="w-64 bg-gray-800 text-white p-4">
                <nav>
                    <Link to="/dashboard">Home</Link>
                </nav>
            </aside>

            {/* Content */}
            <main className="flex-1 p-6">
                <Outlet />
            </main>
        </div>
    );
}
