function handleMutation(mutationsList, observer) {
    // Get elements by class name
    const elements = document.getElementsByClassName("webchat__text-content__markdown");

    // Convert HTMLCollection to an array and iterate over it
    Array.from(elements).forEach(x => {
        console.log(x);
        console.log();
    });
}
