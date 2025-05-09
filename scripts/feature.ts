import { openToast } from "./todo";

export function performToDoAfterMath(event: CustomEvent) {
  if (event.detail.xhr.status === 200) {
    console.log(event.detail.xhr);
    (event.currentTarget as HTMLFormElement)?.reset();
    document.querySelector("dialog-title")!.textContent = "Success";
    document.querySelector("dialog-content")!.textContent =
      "The Todo has been created!";
    (document.querySelector("modalBox") as HTMLDialogElement)!.showModal();
  }
}

export function showEditForm(event: Event) {
  const descElement = event.currentTarget as HTMLSpanElement;
  const form = descElement
    .closest("div")
    ?.querySelector("form") as HTMLFormElement;
  const cancelButton = form.querySelector(
    'button[type="button"]'
  ) as HTMLButtonElement;
  if (!form || !descElement || !cancelButton) {
    openModal(
      "Error",
      "There was an error while trying to open the form. Please try again."
    );
    return;
  }
  descElement.style.display = "none";
  form.style.display = "block";

  cancelButton.onclick = function () {
    form.style.display = "none";
    descElement.style.display = "inline";
  };
}

const modal = document.querySelector(".main-modal") as HTMLDivElement;
const modalContent = document.querySelector("#modal-content") as HTMLDivElement;
const modalTitle = document.querySelector("#modal-title") as HTMLHeadingElement;
const closeButton = document.querySelectorAll(".modal-close");

const modalClose = () => {
  if (!modal) {
    openToast("error", "Modal not found");
    return;
  }
  modal.classList.add("fadeOut");
  setTimeout(() => {
    modal.style.display = "none";
  }, 500);
};

export function openModal(title: string, content: string) {
  if (!modal || !modalContent || !modalTitle) {
    openToast("error", "Modal or modal content/title not found");
    return;
  }

  modal.classList.remove("fadeOut");
  modal.classList.add("fadeIn");
  modal.style.display = "flex";
  modalTitle.innerHTML = title;
  modalContent.innerHTML = content;
}

for (let i = 0; i < closeButton.length; i++) {
  const elements = closeButton[i] as HTMLButtonElement;

  elements.onclick = (e) => modalClose();

  modal.style.display = "none";

  window.onclick = function (event) {
    if (event.target == modal) modalClose();
  };
}
