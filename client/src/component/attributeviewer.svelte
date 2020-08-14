<script lang="ts">
    import {fly} from 'svelte/transition';
    import {getClassFromLevel, copyToClipboard} from '../lib/helper';
    import {createEventDispatcher} from 'svelte';
    import AttrBlock from "./AttrBlock.svelte";
    import {Log} from '../lib/logs';

    export let log: Log;

    const dispatch = createEventDispatcher();

    function closeEvent() {
        dispatch('close', {});
    }
</script>

<style>
    #attrview {
        position: fixed;
        right: 0;
        top: 0;
        height: 100%;
        width: 50%;
        z-index: 1000;
        padding: 1rem;
        background-color: #ffffff;
        font-size: 80%;
        overflow: auto;
    }

    #attrview.debug {
        border-top: 5px solid #9e9e9e;
    }

    #attrview.info {
        border-top: 5px solid #2196F3;
    }

    #attrview.warn {
        border-top: 5px solid #fbc02d;
    }

    #attrview.error {
        border-top: 5px solid #F44336;
    }

    #attrview.critical {
        border-top: 5px solid #000000;
    }

    #attrview .msg {
        border: 1px solid #cccccc;
        padding: 5px;
        overflow-wrap: break-word;
    }

    .clipboard {
        display: none;
        cursor: pointer;
    }

    .clipboard .material-icons {
        font-size: 0.9em;
    }

    #attrview .msg:hover .clipboard{
        display: inline;
    }

    #attrview-header {
        height: 30px;
    }

    .attr-level {
        padding: 3px 5px;
        border-radius: 3px;
        font-weight: bold;
        color: #ffff;
    }

    .control {
        color: #a7a7a7;
    }
</style>

<div id="attrview" class="z-depth-3 {log.level}" transition:fly="{{duration: 300, x: 500}}">
    <div class="row">
        <div id="attrview-header">
            <span class="attr-level {getClassFromLevel(log.level)}">{log.level}</span>

            {log.timestamp}
            <a href="#!" class="right" on:click={closeEvent}>
                <i class="material-icons">close</i>
            </a>
        </div>
        <div class="divider"></div>
    </div>


    <div class="row">
        <div class="col s6 truncate">
            <strong>PROVIDER</strong><br />
            {log.provider}
        </div>

        <div class="col s6 truncate">
            <strong>ID</strong><br />
            {log.id}
        </div>
    </div>

    <div class="row">
        <div class="divider"></div>
    </div>

    <div class="msg yellow lighten-5">
        {log.message}
        <span class="clipboard" on:click={copyToClipboard(log.message)}>
            <i class="tiny material-icons">content_copy</i>
        </span>
    </div>

    {#if log.payload !== null}
        <h6>Attributes</h6>

        <span class="control">{'{'}</span>
        <AttrBlock payload={log.getPayload()} prefix="" />
        <span class="control">{'}'}</span>
    {/if}
</div>
