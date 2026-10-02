<script lang="ts">
    import { onMount } from "svelte";

    "use client"

    import List from "../components/List.svelte";
    import Map from "../components/Map/Map.svelte";
    import Button from "../components/Basic/Button.svelte";
    import Input from "../components/Basic/Input.svelte";
    import { fetch_path, fetch_robot_mode } from "#lib/requests.ts";

    let path = $state([]);
    let robot_mode = $state(0);
    let limit = 4;
    let offset = 0;

    let map_id = $state("");
    let old_map_id = $state("")
    let polling = false;

    let show_list = $state(false);
    
    function toggle_list() {
        show_list = !show_list;
    }

    function clear_canvas() {
        path = [];
    }

    async function poll_path() {
        if(!polling) {
            return;
        }

        const data = await fetch_path(old_map_id, limit, offset)

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

    async function submit_event(event: SubmitEvent) {
        event.preventDefault()
        const data = await fetch_path()

        if(data) {
            path = data
        }
    }

    onMount(async () => {
        robot_mode = await fetch_robot_mode();

        if(robot_mode === 1) {
            toggle_automode();
        }
    })
</script>

<main class="main">
    <div class="topbar">
        <div>
            <Button onclick={toggle_automode}>Start Automode</Button>
            <Button onclick={clear_canvas}>Clear Canvas</Button>
        </div>

        <form id="form" action="/api/coordinates" method="GET" onsubmit={submit_event}>
            <Input type="text" name="mapId" id="" />
            <Button>Draw Map</Button>
        </form>
        
        <Button onclick={toggle_list}>Show Maps</Button>

    </div>
    <Map path={path} robot_mode={robot_mode}></Map>
    <div class="wrapper">
        <div class="wrapper--inner">
            
        </div>
    </div>
    <List robot_mode={robot_mode} visible={show_list}/>
    <div>
        
    </div>
</main>

<style>
    .main {
        position: fixed;
        inset: 0;
        padding: 0px;
        overflow: none;
    }

    .topbar {
        position: absolute;
        display: flex;
        width: 100%;

        justify-content: space-between;
        gap: 16px;
        padding: 8px;

        form {
            display: grid;
            grid-template-columns: 3fr 1fr;
            gap: 8px;

            max-width: 400px;
            width: 400%;
        }
    }
</style>