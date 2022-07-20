import { Requests } from ".";

function Holder () {
    let data = [];
    let updateFuncs = []
    // get data from API URL
    async function refresh() {
        try {
            data = await Requests.getDevices();
        }
        catch (err) {
            console.log(err);
            return
        }
        data = data?.result_list ?? data ?? [];
        // call all update functions with data
        // console.log(data)
        for (const func of updateFuncs) {
            if (func) func(data);
        }

        return data;
    }
    // add update function
    function onUpdate(callback) {
        if (typeof callback !== 'function') return -1
        if (data) callback(data);
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

    function manualUpdate(data) {
        data = data?.result_list ?? data ?? [];
        // call all update functions with data
        for (const func of updateFuncs) {
            if (func) func(data);
        }

        return data;
    }

    const interval = setInterval(refresh, 10000);
    refresh();

    return {
        get: () => data,
        refresh,
        onUpdate,
        removeUpdate,
        manualUpdate,
        interval
    }
}

export default Holder();