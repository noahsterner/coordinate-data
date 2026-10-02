export async function fetch_robot_mode(): Promise<number> {
    const response = await fetch("http://localhost:8080/api/robot/mode");
    const data = await response.json();

    return data.robot_mode == 0 ? 0 : 1;
}

export async function fetch_path(map_id: string = "", limit: number = 0, offset: number = 0): Promise<[]> {
    const form = document.getElementById("form") as HTMLFormElement;
    const submitter = document.querySelector("#form > button") as HTMLElement;
    const formData = new FormData(form, submitter);

    let url = new URL("http://localhost:8080/api/coordinates")

    if (map_id === "") {
        map_id = formData.get("mapId") as string
    }

    url.searchParams.append("mapId", map_id);
    url.searchParams.append("limit", String(limit));
    url.searchParams.append("page", String(offset));

    const response = await fetch(url);
    const data = await response.json();

    return data
}