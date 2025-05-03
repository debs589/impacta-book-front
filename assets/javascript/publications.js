$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.unlike-publication', unlikePublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publication",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Error while creating new publication", "error");
    })
}

function likePublication(event) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const publicationId = elementClicked.closest('div').data('publication-id');

    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/publication/${publicationId}/like`,
        method: "POST"
    }).done(function() {
        const likesCounter = elementClicked.next('span');
        const quantityOfLikes = parseInt(likesCounter.text());

        likesCounter.text(quantityOfLikes + 1);

        elementClicked.addClass('unlike-publication');
        elementClicked.addClass('text-danger');
        elementClicked.removeClass('like-publication');

    }).fail(function() {
        Swal.fire("Ops...", "Error while like the publication!", "error");
    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}

function unlikePublication(event) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const publicationId = elementClicked.closest('div').data('publication-id');

    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/publication/${publicationId}/unlike`,
        method: "POST"
    }).done(function() {
        const likesCounter = elementClicked.next('span');
        const quantityOfLikes = parseInt(likesCounter.text());

        likesCounter.text(quantityOfLikes - 1);

        elementClicked.removeClass('unlike-publication');
        elementClicked.removeClass('text-danger');
        elementClicked.addClass('like-publication');

    }).fail(function() {
        Swal.fire("Ops...", "Error while unlike the publication!", "error");
    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}