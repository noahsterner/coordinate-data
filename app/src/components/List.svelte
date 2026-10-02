<script lang="ts">
    import { onMount } from "svelte";
    
    type Props = {
        robot_mode: number;
    };

    let { robot_mode }: Props = $props();

    let maps = $state<{uuid: string}[]>([])
        
    async function get_maps() {
            const response = await fetch("http://localhost:8080/api/sessions");
            const data = await response.json();

            if(data) {
                maps = data
            }
    }

    onMount(() => {
        get_maps()
    })

    $effect(() => {
        robot_mode;

        if(robot_mode == 0) {
            get_maps()
        }
    })
</script>

<ul>
    {#if maps}
        {#each maps as map}
            <li>{map.uuid}</li>
        {/each}
    {/if}
</ul>

<style>
    li {
        list-style: none;
    }
</style>
