import { Requests } from ".";
import Dummy from "../dummy/data.json";
import { apiStore } from "../stores";

function Holder () {
    let data = [];
    let updateFuncs = []
    const store = apiStore;
    // get data from API URL
    async function refresh() {
        try {
            data = await Requests.getDevices();
        }
        catch (err) {
            console.log(err)
            data = Dummy;
        }
        data = data || Dummy;
        data = data?.result_list ?? data ?? [];
        // call all update functions with data
        for (const func of updateFuncs) {
            if (func) func(data);
        }
        store.data = data;

        return data;
    }
    // add update function
    function onUpdate(callback) {
        if (typeof callback !== 'function') return -1
        callback(data);
        updateFuncs.push(callback);

        return updateFuncs.length - 1;
    }
    // remove update function
    function removeUpdate(callback) {
        if (typeof callback === 'number') {
            updateFuncs.splice(callback, 1);
        }
        else {
            updateFuncs = updateFuncs.filter(func => func !== callback);
        }
    }

    const interval = setInterval(refresh, 10000);
    refresh();

    return {
        get: () => data,
        refresh,
        onUpdate,
        removeUpdate,
        store,
        interval
    }
}

export default Holder();