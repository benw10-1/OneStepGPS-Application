// import { Requests } from ".";
import Dummy from "../dummy/data.json";
import { apiStore } from "../stores";

function Holder () {
    let data = {};
    let udpateFuncs = []
    const store = apiStore;

    async function refresh() {
        // try {
        //     data = await Requests.getInfo();
        // }
        // catch (err) {
        //     data = Dummy;
        // }
        data = Dummy;
        data = data.result_list ?? [];
        for (const func of udpateFuncs) {
            if (func) func(data);
        }
        store.data = data;
    }

    function onUpdate(callback) {
        udpateFuncs.push(callback);
    }

    function removeUpdate(callback) {
        udpateFuncs = udpateFuncs.filter(func => func !== callback);
    }

    refresh();

    const interval = setInterval(refresh, 10000);

    return {
        get: () => data,
        refresh: () => refresh(),
        onUpdate: (callback) => onUpdate(callback),
        removeUpdate,
        store,
        interval
    }
}

export { Holder }

export default Holder();