<script lang="ts">
    import { onMount } from "svelte";
    import { Camera } from "./camera.svelte.ts";
    import type { Vector } from "./map.svelte.ts";

    import CanvasController from "./canvas.svelte.ts";

    type Props = {
        canvas: HTMLCanvasElement;
        robot_mode: number;
        canvas_height: number;
        canvas_width: number;
        path: Vector[];
    };

    let { canvas, robot_mode, path, canvas_height, canvas_width }: Props = $props();

    let canvasController: CanvasController;
    let camera: Camera;

    onMount(() => {
        let drag: any = null
        camera = new Camera(canvas_width/2, canvas_height/2);
        canvasController = new CanvasController(canvas)

        console.log(canvasController);

        canvas.onwheel = (e) => { e.preventDefault(); camera.zoom(e.deltaY); };

        canvas.onmousedown = (e) => {
            drag = {
                offest_x: e.offsetX,
                offest_y: e.offsetY,
                old_camera_x: camera.x,
                old_camera_y: camera.y,
            }
        };

        canvas.onmousemove = (e) => {
            if (drag) {
                let dx = e.offsetX - drag.offest_x;
                let dy = e.offsetY - drag.offest_y;

                camera.x = drag.old_camera_x + dx;
                camera.y = drag.old_camera_y + dy;
            }
        };

        canvas.onmouseup = (e) => { drag = null; };
    })
    
    $effect(() => {
        canvasController.render(camera, path)
    });

    $effect(() => {
        path;

        camera.reset(canvas.width, canvas.height)
    })
</script>

<canvas bind:this={canvas}>

</canvas>