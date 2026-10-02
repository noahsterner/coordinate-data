<script lang="ts">
    import Canvas from "./Canvas.svelte";
    import type { Vector } from "./map.svelte.ts";

    const canvas_width: number = 800;
    const canvas_height: number = 600;

    type Props = {
        robot_mode: number;
        path: Vector[];
    };

    let { path, robot_mode }: Props = $props();
</script>

<Canvas canvas_width={canvas_width} canvas_height={canvas_height} path={path} robot_mode={robot_mode} />