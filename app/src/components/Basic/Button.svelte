<script lang="ts">
    import type { Snippet } from "svelte";
    import type { HTMLButtonAttributes } from "svelte/elements";

    type Variant = "filled" | "outlined";
 

    type ButtonProps = HTMLButtonAttributes & {
        children?: Snippet;
        variant?: Variant;
    };
 
    let { children, variant = "filled", ...rest }: ButtonProps = $props()
</script>

<button
    class={`button button-${variant}`}
    {...rest}
>
    {@render children?.()}
</button>

<style>
    button {
        padding: 8px;
        border: none;
        border-radius: 5px;
        color: var(--on-primary-500);
    }

    .button-filled {
        background-color: var(--primary-500);
    }
</style>