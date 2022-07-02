import { Requests } from ".";

function Holder () {
    let data = {};
    let udpateFunc = () => {};

    async function refresh() {
        data = await Requests.getInfo();
        udpateFunc(data);
    }

    function onUpdate(callback) {
        udpateFunc = callback;
    }

    refresh();

    return {
        get: () => data,
        refresh: () => refresh(),
        onUpdate: (callback) => onUpdate(callback)
    }
}

export { Holder }

export default Holder();