<script lang="ts">
    import type { Vector } from "./map.svelte.ts";

    type Props = {
        canvas_height: number;
        canvas_width: number;
        path: Vector[];
        origo: Vector;
    };

    let { path, canvas_height, canvas_width, origo = $bindable() }: Props = $props();

    let canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D;

    const scale_speed = 0.02 * 1;
    const max_zoom = 1;
    const min_zoom = 0.01;

    let scale: number = $state(0.15);
    let mouse_down: boolean = false;

    let down_offset: Vector = { x: origo.x, y: origo.y };
    let origo_down: Vector = { x: origo.x, y: origo.y };

    function draw_px(
        ctx: CanvasRenderingContext2D,
        vector: Vector,
        vector_next: Vector,
    ) {
        ctx.moveTo(
            origo.x + vector_next.x * scale,
            origo.y - vector_next.y * scale,
        );
        ctx.lineTo(origo.x + vector.x * scale, origo.y - vector.y * scale);
    }

    function start_drawing() {
        ctx.clearRect(0, 0, canvas_width, canvas_height);

        for (let i = 0; i < path.length - 1; ++i) {
            let next = i + 1;

            draw_px(ctx, path[i], path[next]);
        }

        ctx.stroke();
    }

    $effect(() => {
        ctx = canvas.getContext("2d") as CanvasRenderingContext2D;

        canvas.height = canvas_height;
        canvas.width = canvas_width;

        canvas.onwheel = (e) => {
            e.preventDefault();
            if (max_zoom > scale && e.deltaY < 0) {
                scale += scale_speed;
                if (scale > max_zoom) scale = max_zoom;
            } else if (min_zoom < scale && e.deltaY > 0) {
                scale -= scale_speed;
                if (scale < min_zoom) scale = min_zoom;
            }
        };

        canvas.onmousedown = (e) => {
            e.preventDefault();
            down_offset.x = e.offsetX;
            down_offset.y = e.offsetY;
            origo_down = { x: origo.x, y: origo.y };

            mouse_down = true;
        };

        canvas.onmousemove = (e) => {
            e.preventDefault();
            if (mouse_down) {
                let dx = e.offsetX - down_offset.x;
                let dy = e.offsetY - down_offset.y;

                origo.x = origo_down.x + dx;
                origo.y = origo_down.y + dy;

                console.log(e);
            }
        };

        canvas.onmouseup = (e) => {
            e.preventDefault();
            mouse_down = false;
        };

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
