// src/App.tsx

import { useMetrics } from "./useMetrics";

function App() {
  const metrics = useMetrics();

  return (
    <div className="p-6">
      <h1 className="text-xl font-bold">Agent Metrics ({metrics.length})</h1>
      <div className="mt-4 space-y-2 flex gap-3 flex-wrap">
        {metrics.map((m, i) => (
          <div key={i} className="bg-violet-50 font-bold text-violet-900 border border-via-violet-200 rounded p-2 w-10 h-10 grid place-items-center">
            {m.cpu}
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
