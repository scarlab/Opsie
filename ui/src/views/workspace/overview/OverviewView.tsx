import NodesOverview from "./NodesOverview";
import OverviewCharts from "./OverviewCharts";
import ProjectsOverview from "./ProjectsOverview";
import ResourceOverview from "./ResourceOverview";

export default function OverviewView() {
    return (
        <div className="">
            <OverviewCharts />
            <NodesOverview />
            <ProjectsOverview />
            <ResourceOverview />
        </div>
    )
}
