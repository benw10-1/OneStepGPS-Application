import { Requests } from ".";
import Dummy from "../dummy/data.json";
import { apiStore } from "../stores";

function Holder () {
    let data = {};
    let updateFuncs = []
    const store = apiStore;
    // get data from API URL
    async function refresh() {
        try {
            data = await Requests.getAPIData();
        }
        catch (err) {
            data = Dummy;
        }
        data = data?.result_list ?? [];
        for (const func of updateFuncs) {
            if (func) func(data);
        }
        store.data = data;

        return data;
    }

    function onUpdate(callback) {
        updateFuncs.push(callback);
        return updateFuncs.length - 1;
    }

    function removeUpdate(callback) {
        if (typeof callback === 'number') {
            updateFuncs.splice(callback, 1);
        }
        else {
            updateFuncs = updateFuncs.filter(func => func !== callback);
        }
    }

    refresh();

    // const interval = setInterval(refresh, 10000);

    return {
        get: () => data,
        refresh,
        onUpdate,
        removeUpdate,
        store,
        // interval
    }
}

export { Holder }

export default Holder();