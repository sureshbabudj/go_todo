import { openModal } from "./feature";

export function newTodo(event: CustomEvent) {
  if (event.detail.xhr.status === 200) {
    (event.currentTarget as HTMLFormElement)?.reset();
    updateDialog("success", "The Todo has been created!");
    refreshTodo();
  }
}

export function deleteTodo(event: CustomEvent) {
  if (event.detail.xhr.status === 200) {
    updateDialog("success", "The Todo has been deleted!");
    refreshTodo();
  }
}

export function updateTodo(event: CustomEvent) {
  if (event.detail.xhr.status === 200) {
    updateDialog("success", "The Todo has been updated!");
    refreshTodo();
  }
}

export function updateDialog(title: string, content: string) {
  // openModal(title, content);
  openToast("success", content);
}

export function refreshTodo() {
  const refreshTodos = document.querySelector(
    "#refreshTodos"
  ) as HTMLButtonElement;
  if (!refreshTodos) {
    openModal(
      "Error",
      "There was an error while trying to refresh the todos. Please try again."
    );
    return;
  }
  refreshTodos.click();
}

export function openToast(
  type: string,
  content: string,
  duration: number = 3000
) {
  const types = ["info", "success", "warning", "error"];

  if (!types.includes(type)) {
    openModal(
      "Error",
      "There was an error while trying to open the toast. Please try again."
    );
    return;
  }

  const toast = document.createElement("div");
  toast.className = "toast";
  toast.innerHTML = `
  <div class="toast">
    <div class="alert alert-${type || "info"}">
      <span id="toast-content">${content}</span>
    </div>
  </div>
  `;
  document.body.appendChild(toast);
  setTimeout(() => {
    toast.remove();
  }, duration);
}
