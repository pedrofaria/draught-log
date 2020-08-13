<script>
    import {copyToClipboard} from '../lib/helper';
    import AttrBlock from "./AttrBlock.svelte";

    export let prefix;
    export let key;
    export let value;

    let hide = false;
    let computedKey = (prefix === '') ? key : prefix + '.' + key;
</script>

<style>
    tr {
        border: 0;
    }

    .attr-item:hover {
        background-color: #fafaea;
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

    .material-icons {
        font-size: 0.9em;
    }

    .clipboard {
        display: none;
        cursor: pointer;
    }

    .attr-value:hover .clipboard {
        display: inline;
    }

    .attr-sub {
        padding: 0 0 0 20px;
    }

    .control {
        color: #a7a7a7;
    }
</style>

{#if typeof value !== "object"}
<tr class="attr-item">
    <td class="attr-key raw">{key}</td>
    <td class="attr-value">
        {value}
        <span class="clipboard" on:click={copyToClipboard(value)}>
            <i class="tiny material-icons">content_copy</i>
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
    <td colspan="2" class="attr-sub" >
        <AttrBlock prefix={computedKey} payload={value} />
    </td>
</tr>
<tr class:hide>
    <td colspan="2" class="raw">
        <span class="control">{'}'}</span>
    </td>
</tr>
{/if}