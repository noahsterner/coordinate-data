<script lang="ts">
    import { onMount } from "svelte";
    
    type Props = {
        visible: boolean;
        robot_mode: number;
    };

    let { robot_mode, visible }: Props = $props();

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

{#if visible }
<ul>
    {#if maps}
        {#each maps as map}
            <li>{map.uuid}</li>
        {/each}
    {/if}
</ul>
{/if}

<style>
    ul {
        position: absolute;
        top: 40px; right: 0;

        width: fit-content;
        height: 90vh;
        overflow: scroll;
    }

    li {
        list-style: none;
        padding: 4px;

        &:nth-child(odd) {
            background-color: lightgray;
        }
    }
</style>
