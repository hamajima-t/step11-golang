(() => {
  let socket = null;
  
  document.addEventListener('DOMContentLoaded', function() {
    socket = new WebSocket("ws://127.0.0.1:8080/ws");

    socket.onopen = () => {
      console.log("Successfully connected!");
    };
  });

  // 削除
  $('#items').on('click', '.delete_item', function() {
    let id = $(this).parents('tr').data('id');
    let jsonData = {};
    jsonData['action'] = 'delete';
    jsonData['id'] = id;
    socket.send(JSON.stringify(jsonData));

    $(this).parents('tr').fadeOut();
  });


})();