<script>
    import AttrItem from "./AttrItem.svelte";

    export let payload;
    export let prefix;
</script>

<table>
    <tbody>
    {#each Object.keys(payload) as attr}
        <AttrItem prefix={prefix} key={attr} value={payload[attr]} />
    {/each}
    </tbody>
</table>
