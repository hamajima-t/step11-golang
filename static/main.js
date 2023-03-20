(() => {
  'use strict';

  document.addEventListener('DOMContentLoaded', () => {
    const socket = setupWebSocket();

    // 削除
    const items = document.querySelector('#items');
    items.addEventListener('click', (e) => {
      deleteItem(e.target, socket);
    });
  });

  function setupWebSocket() {
    // WebSocketオブジェクトの作成
    const socket = new WebSocket("ws://127.0.0.1:8080/ws");
    // 接続を確立
    socket.addEventListener('open', () => {
      console.log("Successfully connected!");
    });

    return socket;
  }

  function deleteItem(target, socket) {
    const row = target.closest('tr');
    const id = parseInt(row.dataset.id);
    const jsonData = {
      action: 'delete',
      id: id
    };
    socket.send(JSON.stringify(jsonData));
    row.style.display = 'none';
  }

})();