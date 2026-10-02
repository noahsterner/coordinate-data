<script lang="ts">
    import type { Snippet } from "svelte";
    import type { HTMLInputAttributes } from "svelte/elements";

    type Variant = "filled" | "outlined";
 

    type ButtonProps = HTMLInputAttributes & {
        variant?: Variant;
    };
 
    let { children, variant = "filled", ...rest }: ButtonProps = $props()
</script>

<input
    class={`input input-${variant}`}
    {...rest}
/>

<style>
    input {
        outline: none;
        padding: 8px;
        border: 2px solid var(--surface-900);
        border-radius: 5px;
        color: var(--on-surface-500);
    }

    .input-filled {
        background-color: var(--surface-800);
    }
</style>