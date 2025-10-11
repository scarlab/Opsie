import OverviewCharts from "./OverviewCharts";
import ProjectsOverview from "./ProjectsOverview";
import ResourceOverview from "./ResourceOverview";

export default function OverviewView() {
    return (
        <div className="">
            <OverviewCharts />
            <ProjectsOverview />
            <ResourceOverview />
        </div>
    )
}
