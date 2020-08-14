<script lang="ts">
    import AttrItem from "./AttrItem.svelte";

    export let prefix;
    export let payload: object;
</script>

<table>
    <tbody>
    {#each Object.keys(payload) as attr}
        <AttrItem prefix={prefix} key={attr} value={payload[attr]} />
    {/each}
    </tbody>
</table>
