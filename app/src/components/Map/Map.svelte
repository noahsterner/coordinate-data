<script lang="ts">
    import { onMount } from "svelte";
    import Canvas from "./Canvas.svelte";
    import CanvasViewModel from "./canvas.svelte.ts";
    import type { Vector } from "./map.svelte.ts";

    const canvas_width: number = 800;
    const canvas_height: number = 600;

    type Props = {
        robot_mode: number;
        path: Vector[];
    };

    let { path, robot_mode }: Props = $props();

    let canvas: HTMLCanvasElement = {} as HTMLCanvasElement;
</script>

<Canvas canvas={canvas} canvas_width={canvas_width} canvas_height={canvas_height} path={path} robot_mode={robot_mode} />