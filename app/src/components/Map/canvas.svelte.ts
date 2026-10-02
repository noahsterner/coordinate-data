import type { Camera } from "./camera.svelte";
import type { Vector } from "./map.svelte";

export default class CanvasController {
    private canvas: HTMLCanvasElement
    private ctx: CanvasRenderingContext2D ;

    constructor(canvas: HTMLCanvasElement) {
        this.canvas = canvas;
        const ctx = canvas.getContext("2d") ;

        if(!ctx) {
            throw new Error("Could not get canvas context")
        }

        this.ctx = ctx;
    }

    render(camera: Camera, path: Vector[]) {        
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
        this.ctx.beginPath()
        for (let i = 0; i < path.length - 1; ++i) {
            let next = i + 1;

            this._draw_px(camera, path[i], path[next]);
        }

        this.ctx.stroke();
    }

    private _draw_px(
        camera: Camera,
        vector: Vector,
        vector_next: Vector,
    ) {
        this.ctx.moveTo(
            camera.x + vector_next.x * camera.scale,
            camera.y - vector_next.y * camera.scale,
        );
        this.ctx.lineTo(camera.x + vector.x * camera.scale, camera.y - vector.y * camera.scale);
    }

}