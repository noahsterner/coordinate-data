<script>
    import { onMount } from "svelte";

    let maps = $state([])

    onMount(() => {
        async function get_maps() {
            const response = await fetch("http://localhost:8080/api/sessions");
            const data = await response.json();
       	
		    console.log(data)
            maps = data
        }

        
        get_maps()
    })


    $effect(() => {
        maps;
    })
</script>

<ul>
    {#each maps as map}
        <li>{map.uuid}</li>
    {/each}
</ul>

<style>
    li {
        list-style: none;
    }
</style>
