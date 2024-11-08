import { s as subscribe } from './utils-Cb-NyS5y.js';
import { c as create_ssr_component, e as escape } from './ssr-CwU1ZNF3.js';
import { p as page } from './stores-ES20sfbp.js';
import './client-BUusD8wq.js';
import './exports-BGi7-Rnc.js';

const Error = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $page, $$unsubscribe_page;
  $$unsubscribe_page = subscribe(page, (value) => $page = value);
  $$unsubscribe_page();
  return `<h1>${escape($page.status)}</h1> <p>${escape($page.error?.message)}</p>`;
});

export { Error as default };
//# sourceMappingURL=error.svelte-CIpbHqRg.js.map
