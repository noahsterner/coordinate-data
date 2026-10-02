<script lang="ts">
    import { onMount } from "svelte";

    "use client"

    import List from "../components/List.svelte";
    import Map from "../components/Map/Map.svelte";

    let path = $state([]);
    let robot_mode = $state(0);
    let limit = 4;
    let offset = 0;

    let map_id = $state("");
    let old_map_id = $state("")
    let polling = false;

    async function poll_path() {
        if(!polling) {
            return;
        }

        const data = await get_path(old_map_id, limit, offset)

        if(data.length > 0) {
            offset += data.length;
            path.push(...data as []);
        } else if(data.length <= 0 && robot_mode === 0) {
            polling = false;
            return;
        }

        if(!polling) {
            return;
        }

        setTimeout(poll_path, 150)
    }

    async function toggle_automode() {
        robot_mode = robot_mode === 0 ? 1 : 0;

        const response = await fetch("http://localhost:8080/api/robot/mode", {
            method: "POST",
            headers: { "Content-Type": "application/json"},
            body: JSON.stringify({robot_mode: robot_mode}),
        });

        const data = await response.json();
        map_id = data.map_id;

        if(robot_mode === 1) {
            path = [];
            offset = 0;
            polling = true;
            old_map_id = map_id;
            setTimeout(poll_path, 200)
        }
    }

    async function get_path(map_id: string = "", limit: number = 0, offset: number = 0) {
        const form = document.getElementById("form") as HTMLFormElement;
        const submitter = document.querySelector("#form > button") as HTMLElement;
        const formData = new FormData(form, submitter);

        let url = new URL("http://localhost:8080/api/coordinates")
        
        if(map_id === "") {
            map_id = formData.get("mapId") as string
        }

        url.searchParams.append("mapId", map_id);
        url.searchParams.append("limit", String(limit));
        url.searchParams.append("page", String(offset));

        const response = await fetch(url);
        const data = await response.json();

        return data
    }

    async function submit_event(event: SubmitEvent) {
        event.preventDefault()
        const data = await get_path()

        if(data) {
            path = data
        }
    }

    async function fetch_robot_mode() {
        const response = await fetch("http://localhost:8080/api/robot/mode");
        const data = await response.json();

        robot_mode = data.robot_mode;
    }

    onMount(async () => {
        await fetch_robot_mode();

        if(robot_mode === 1) {
            toggle_automode();
        }
    })
</script>

<main>

    <div class="wrapper">
        <Map path={path} robot_mode={robot_mode}></Map>

        <div class="wrapper--inner">
            <form id="form" action="/api/coordinates" method="GET" onsubmit={submit_event}>
                <input type="text" name="mapId" id="">
                <button>Draw Map</button>
            </form>
            <button onclick={toggle_automode}>Start Automode</button>
        </div>
    </div>
    <List robot_mode={robot_mode}/>
</main>

<style>
    main {
        width: 800px;
    }

    .wrapper--inner {
        display: flex;
        gap: 32px;

        form { 
            flex-grow: 10;
            display: flex;
            gap: 4px;
            max-width: 100%;
            
            input {
                flex-grow: 10;
            }

            button {
                flex-grow: 1;
            }
        }
    }

</style>