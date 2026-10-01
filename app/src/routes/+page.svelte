<script lang="ts">
    "use client"

    import List from "../components/List.svelte";
    import Map from "../components/Map/Map.svelte";

    let path = $state([])

    async function get_path(limit: number = 0, offset: number = 0) {
        const form = document.getElementById("form") as HTMLFormElement;
        const submitter = document.querySelector("button[value=save]") as HTMLElement;
        const formData = new FormData(form, submitter);

        let url = new URL("http://localhost:8080/api/coordinates")
        
        const mapId = formData.get("mapId")
        if(!mapId) return

        url.searchParams.append("mapId", String(mapId));
        url.searchParams.append("limit", String(limit));
        url.searchParams.append("offset", String(offset));

        const response = await fetch(url);
        const data = await response.json();
            
        path = data
        //console.log(path)
    }

    function submit_event(event: SubmitEvent) {
        event.preventDefault()
        get_path()
    }
</script>

<main>
    <div class="wrapper">
        <Map path={path}></Map>
        <div class="wrapper--inner">
            <form id="form" action="/api/coordinates" method="GET" onsubmit={submit_event}>
                <input type="text" name="mapId" id="">
                <button>Draw Map</button>
            </form>
        </div>
    </div>
    <List />
</main>

<style>
    main {
        width: 800px;
    }

    .wrapper--inner {
        form { 
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