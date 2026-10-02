export class Camera {
    x: number;
    y: number;
    scale: number;
    min_zoom: number;
    max_zoom: number;
    scale_speed: number;

    constructor(
        canvas_width: number,
        canvas_height: number,
        scale: number = 0.15,
        min_zoom: number = 0.01,
        max_zoom: number = 1
    ) {
        this.x = $state(canvas_width);
        this.y = $state(canvas_height);
        this.scale = $state(scale);
        this.min_zoom = min_zoom;
        this.max_zoom = max_zoom;
        this.scale_speed = 0.01;
    }

    reset(canvas_width: number, canvas_height: number) {
        this.x = canvas_width / 2;
        this.y = canvas_height / 2;
        this.scale = 0.15
    }

    zoom(delta_y: number) {
        if (this.max_zoom > this.scale && delta_y < 0) {
            this.scale += this.scale_speed;
            if (this.scale_speed > this.max_zoom) this.scale = this.max_zoom;
        } else if (this.min_zoom < this.scale && delta_y > 0) {
            this.scale -= this.scale_speed;
            if (this.scale < this.min_zoom) this.scale = this.min_zoom;
        }
    }
}