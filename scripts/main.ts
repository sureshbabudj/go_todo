import htmx from "htmx.org";
import * as featureFunctions from "./feature";
import * as todoFunctions from "./todo";

declare global {
  interface Window {
    htmx: typeof htmx;
  }
}

[...Object.entries(todoFunctions), ...Object.entries(featureFunctions)].forEach(
  ([key, value]) => {
    if (typeof value === "function") {
      (window as any)[key] = value;
    }
  }
);

// Attach the functions to the global `window` object
window.htmx = htmx;
htmx.defineExtension("addhtmlheader", {
  onEvent: function (name: string, evt: Event) {
    if (name === "htmx:configRequest") {
      (evt as CustomEvent).detail.headers["Accept"] = "text/html";
    }
    return true;
  },
});

htmx.defineExtension("submitjson", {
  onEvent: function (name: string, evt: Event) {
    if (name === "htmx:configRequest") {
      (evt as CustomEvent).detail.headers["Content-Type"] = "application/json";
      (evt as CustomEvent).detail.headers["Accept"] = "text/html";
    }
    return true;
  },
  encodeParameters: function (
    xhr: XMLHttpRequest,
    parameters: FormData,
    elt: Node
  ): string {
    xhr.overrideMimeType("text/json");
    const jsonParameters: Record<string, boolean | string> = {};
    parameters.forEach((value, key) => {
      if (value === "on") {
        jsonParameters[key] = true;
      } else if (value === "off") {
        jsonParameters[key] = false;
      } else {
        jsonParameters[key] = value instanceof File ? value.name : value;
      }
    });
    const body = {
      ...jsonParameters,
    };
    return JSON.stringify(body);
  },
});
