<script lang="ts">
    import { onMount } from "svelte";
    import { Camera } from "./camera.svelte.ts";
    import type { Vector } from "./map.svelte.ts";

    type Props = {
        robot_mode: number;
        canvas_height: number;
        canvas_width: number;
        path: Vector[];
    };

    let { robot_mode, path, canvas_height, canvas_width }: Props = $props();

    let canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D;
    let camera: Camera;

    function draw_px(
        ctx: CanvasRenderingContext2D,
        vector: Vector,
        vector_next: Vector,
    ) {
        ctx.moveTo(
            camera.x + vector_next.x * camera.scale,
            camera.y - vector_next.y * camera.scale,
        );
        ctx.lineTo(camera.x + vector.x * camera.scale, camera.y - vector.y * camera.scale);
    }

    function start_drawing() {
        ctx.clearRect(0, 0, canvas_width, canvas_height);
        ctx.beginPath()
        for (let i = 0; i < path.length - 1; ++i) {
            let next = i + 1;

            draw_px(ctx, path[i], path[next]);
        }

        ctx.stroke();
    }

    onMount(() => {
        let drag: any = null
        camera = new Camera(canvas_width/2, canvas_height/2);


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
        ctx = canvas.getContext("2d") as CanvasRenderingContext2D;

        canvas.height = canvas_height;
        canvas.width = canvas_width;

        start_drawing()
    });
</script>

<canvas bind:this={canvas}>

</canvas>

<style>
    canvas {
        border: 1px solid black;
    }
</style>
