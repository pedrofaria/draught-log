<script lang="ts">
    import Attributeviewer from './attributeviewer.svelte'
    import { Log } from '../lib/logs';
    import {getClassFromLevel} from '../lib/helper';

    let MAX_ITEMS = 300;
    let items: Log[] = [];
    let selectedItem: Log | null = null;
    let totalLogs = 0;
    let searchTerm: string = "";
    let extraLogCols: string[] = [];
    let skipNonProcessed = false;

    function getPreparedItems(items: Log[], search: string, skip: bool): Log[] {
        // filter
        let list = items.filter(item => {
            if (skip && item.payload == null) {
               return false;
            }
            return item.message.toLowerCase().indexOf(search.toLowerCase()) !== -1
        });

        // sort
        const sortField = 'timestamp';
        const sortAscending = false;

        function sortDirection(a, b, asc) {
            return asc ? a > b : b > a;
        }

        function comp(a, b) {
            if (sortDirection(a[sortField], b[sortField], sortAscending)) { return 1; }
            if (sortDirection(b[sortField], a[sortField], sortAscending)) { return -1; }
            return 0;
        }
        return list.sort(comp);
    }

    $: filteredList = getPreparedItems(items, searchTerm, skipNonProcessed);
    $: extraColumns = extraLogCols;

    const formatter = new Intl.DateTimeFormat('en', {
        day: "2-digit",
        month: "short",
        hour12: false,
        hour: 'numeric',
        minute: '2-digit',
        second: 'numeric',
        fractionalSecondDigits: 1
    });

    function formatDate(ts) {
        return formatter.format(ts);
    }

    function addItem(item: object) {
        let log = new Log(
            item.id,
            item.provider,
            item.raw_log,
            item.message,
            item.level,
            item.timestamp,
            item.payload
        );

        items.unshift(log);
        totalLogs += 1;
        if (items.length > MAX_ITEMS) {
            items.splice(MAX_ITEMS - 1, items.length - MAX_ITEMS);
        }

        items = items;
    }

    function clearLogs() {
        items = [];
        totalLogs = 0;
    }

    function selectItem(item: Log) {
        if (selectedItem !== null && item.id == selectedItem.id) {
            selectedItem = null;
            return;
        }
        document.Item = item;

        selectedItem = item;
    }

    function handleClose() {
        selectedItem = null;
    }

    function handleAddColumn(event) {
        if (extraLogCols.indexOf(event.detail.column) !== -1) {
            removeColumn(event.detail.column);
            return
        }

        extraLogCols.push(event.detail.column);
        extraLogCols = extraLogCols;
    }

    function removeColumn(column) {
        extraLogCols.splice(extraLogCols.indexOf(column), 1);
        extraLogCols = extraLogCols;
    }

    // Create WebSocket connection.
    const socket = new WebSocket('ws://' + location.host + '/api/stream');

    // Listen for messages
    socket.addEventListener('message', function (event) {
        let item = JSON.parse(event.data);
        addItem(item);
    });

</script>

<style>
    #searchbox .input-field, #searchbox .input-field input {
        margin: 0;
    }

    #searchbox .card-content .row:first-of-type {
        margin-bottom: 0;
    }

    #logpanel {
        font-size: 80%;
    }

    #logview {
        width: 100%;
        max-width: 100%;
        max-height: 100%;
        border-collapse: collapse;
        margin: 0;
        padding: 0;
    }

    #logview tr.selected {
        background-color: #eed1b1;
    }

    #logview td, #logview th {
        padding: 0.3rem;
    }

    .header-level {
        width: 1.5rem;
    }

    .header-timestamp {
        min-width: 9rem;
    }

    .header-extra {
        min-width: 9rem;
        max-width: 13rem;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .header-message {
        min-width: 300px;
    }

    .header-last {
        max-width: none;
        width: 100%;
    }

    .header {
        text-align: left;
        max-width: 46px;
        position: relative;
        padding: 0;
        border-top: 1px solid var(--ui-border);
    }

    .field {
        white-space: nowrap;
    }

    .field-extra {
        max-width: 25rem;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .field-message {
        overflow: hidden;
    }

    .level {
        width: 0.5rem;
        height: 1rem;
    }
</style>

<div class="row">
    <div class="col s12">
        <div id="searchbox" class="card">
            <div class="card-content">
                <div class="row">
                    <div class="col s12">
                        <div class="input-field">
                            <input id="filterInput" type="text" class="validate" bind:value={searchTerm}>
                            <label for="filterInput">Filter</label>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col s12">
                        <div class="input-field">
                            <label>
                                <input type="checkbox" bind:checked={skipNonProcessed} />
                                <span>Skip logs without attributes</span>
                            </label>
                        </div>
                    </div>
                </div>
            </div>
            <div class="card-action">
                <div class="right">
                    <a href="#!" on:click={clearLogs}>Clear Logs</a>
                </div>
                <div>
                    Total logs: {totalLogs}
                </div>
            </div>
        </div>

    </div>
</div>

<div id="logpanel" class="row">
    <div class="col s12">
        <table id="logview" class="highlight">
            <thead>
            <tr>
                <th class="header header-level"></th>
                <th class="header header-timestamp">Timestamp</th>
                <th class="header header-provider">Provider</th>
                {#each extraColumns as extraCol}
                    <th class="header header-extra" on:click={removeColumn(extraCol)}>{extraCol}</th>
                {/each}
                <th class="header header-message header-last">Message</th>
            </tr>
            </thead>

            <tbody>
            {#each filteredList as item (item.id)}
                <tr on:click={selectItem(item)} class:selected={selectedItem === item}>
                    <td>
                        <div class="level {getClassFromLevel(item.level)}"></div>
                    </td>
                    <td class="field field-date">{formatDate(item.timestamp)}</td>
                    <td class="field">{item.provider}</td>
                    {#each extraColumns as extraCol}
                    <td class="field field-extra"
                        title="{item.getPayloadAttr(extraCol)}"
                    >{item.getPayloadAttr(extraCol)}</td>
                    {/each}
                    <td class="field field-message">{item.message}</td>
                </tr>
            {/each}
            </tbody>
        </table>

    </div>
</div>

{#if selectedItem !== null}
    <Attributeviewer on:close={handleClose} on:addcolumn={handleAddColumn} log={selectedItem} />
{/if}
