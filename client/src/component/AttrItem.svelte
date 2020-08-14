<script lang="ts">
    import AttrBlock from "./AttrBlock.svelte";
    import {createEventDispatcher} from 'svelte';
    import {copyToClipboard} from '../lib/helper';

    const dispatch = createEventDispatcher();

    function addColumn(e, col) {
        function check(elem): bool {
            const blacklist: string[] = ["clipboard", "clipboard-icon"];

            return blacklist.filter(value => elem.classList.contains(value)).length > 0
        }

        if (check(e.target) || check(e.target.parentElement)) {
            return
        }

        dispatch('addcolumn', {column: col});
    }

    export let prefix;
    export let key;
    export let value;

    let hide = false;
    $: computedKey = (prefix === '') ? key : prefix + '.' + key;
</script>

<style>
    tr {
        border: 0;
    }

    .attr-item:hover {
        background-color: #fafaea;
        cursor: pointer;
    }

    td {
        padding: 1px 3px;
    }

    .attr-key {
        color: #fb8c00 !important;
        vertical-align: top;
        white-space: nowrap;
    }

    .raw {
        padding-left: 18px;
    }

    .attr-value {
        white-space: pre-wrap;
        word-break: break-word;
        width: 100%;
    }

    .attr-sub {
        padding: 0 0 0 20px;
    }

    .control {
        color: #a7a7a7;
    }

    .material-icons {
        font-size: 0.9em;
    }

    .clipboard {
        display: none;
        cursor: pointer;
    }
</style>

{#if typeof value !== "object"}
<tr class="attr-item" data-attr-key="{computedKey}" on:click={(e) => addColumn(e, computedKey)}>
    <td class="attr-key raw">{key}</td>
    <td class="attr-value">
        {value}
        <span class="clipboard" on:click={() => copyToClipboard(value)}>
            <i class="tiny material-icons clipboard-icon">content_copy</i>
        </span>
    </td>
</tr>

{:else}
<tr class="attr-item" on:click={() => hide = !hide}>
    <td class="attr-key">
        {#if hide}
            <i class="tiny material-icons">expand_more</i>
            {key}
            <span class="control">{'{...}'}</span>
        {:else}
            <i class="tiny material-icons">expand_less</i>
            {key}
            <span class="control">{'{'}</span>
        {/if}
    </td>
    <td class="attr-value"></td>
</tr>
<tr class:hide>
    <td colspan="2" class="attr-sub" data-attr-key="{computedKey}">
        <AttrBlock prefix={computedKey} payload={value} on:addcolumn />
    </td>
</tr>
<tr class:hide>
    <td colspan="2" class="raw">
        <span class="control">{'}'}</span>
    </td>
</tr>
{/if}