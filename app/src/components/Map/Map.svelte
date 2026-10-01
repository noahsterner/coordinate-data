<script lang="ts">
    import Canvas from "./Canvas.svelte";
    import type { Vector } from "./map.svelte.ts";

    const canvas_width: number = 800;
    const canvas_height: number = 600;

    type Props = {
        path: Vector[];
    };

    let { path }: Props = $props();
    
    let origo: Vector = $state({ x: 0, y: 0 })

    function reset_origo() {
        origo.x = canvas_width / 2;
        origo.y = canvas_height / 2;
    }

    $effect(() => {
        path;

        reset_origo()
    })
</script>

<Canvas bind:origo={origo} canvas_width={canvas_width} canvas_height={canvas_height} path={path} />