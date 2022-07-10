function formatSeconds(seconds) {
    seconds = Math.floor(seconds);

    const days = Math.floor(seconds / (60 * 60 * 24));
    const hours = Math.floor((seconds % (60 * 60 * 24)) / (60 * 60));
    const minutes = Math.floor((seconds % (60 * 60)) / 60);
    const seconds_ = Math.floor(seconds % 60);

    const days_str = days > 0 ? `${days}d ` : '';
    const hours_str = hours > 0 ? `${hours}h ` : '';
    const minutes_str = minutes > 0 ? `${minutes}m ` : '';
    const seconds_str = seconds_ > 0 ? `${seconds_}s` : '';

    return `${days_str}${hours_str}${minutes_str}${seconds_str}`;
}

export default formatSeconds;