import { s as subscribe } from './utils-Cb-NyS5y.js';
import { c as create_ssr_component, v as validate_component } from './ssr-CwU1ZNF3.js';
import './client-BUusD8wq.js';
import { T as Tabs, a as TabItem } from './Tabs-hpKObKEs.js';
import { d as derived, w as writable } from './index2-CpnJNRxb.js';
import { p as page } from './stores-ES20sfbp.js';
import './exports-BGi7-Rnc.js';
import './bundle-mjs-BLGuyV0D.js';

const createActiveTabStore = () => {
  const { subscribe: subscribe2, set } = writable("teams");
  return {
    subscribe: subscribe2,
    set: (value) => {
      set(value);
    },
    init: () => {
    }
  };
};
const activeTab = createActiveTabStore();
const routeTab = derived(page, ($page) => {
  const path = $page.url.pathname;
  const tab = path.split("/").pop();
  if (["teams", "pools", "seeding", "bracket", "results"].includes(tab)) {
    activeTab.set(tab);
    return tab;
  }
  return "teams";
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $routeTab, $$unsubscribe_routeTab;
  $$unsubscribe_routeTab = subscribe(routeTab, (value) => $routeTab = value);
  $$unsubscribe_routeTab();
  return `${validate_component(Tabs, "Tabs").$$render(
    $$result,
    {
      class: "font-heading ml-2",
      contentClass: "",
      activeClasses: "p-4 text-theme border-b-2 border-theme"
    },
    {},
    {
      default: () => {
        return `${validate_component(TabItem, "TabItem").$$render(
          $$result,
          {
            open: $routeTab === "teams",
            title: "Teams"
          },
          {},
          {}
        )} ${validate_component(TabItem, "TabItem").$$render(
          $$result,
          {
            open: $routeTab === "pools",
            title: "Pools"
          },
          {},
          {}
        )} ${validate_component(TabItem, "TabItem").$$render(
          $$result,
          {
            open: $routeTab === "seeding",
            title: "Seeding"
          },
          {},
          {}
        )} ${validate_component(TabItem, "TabItem").$$render(
          $$result,
          {
            open: $routeTab === "bracket",
            title: "Bracket"
          },
          {},
          {}
        )} ${validate_component(TabItem, "TabItem").$$render(
          $$result,
          {
            open: $routeTab === "results",
            title: "Results"
          },
          {},
          {}
        )}`;
      }
    }
  )} <br> ${slots.default ? slots.default({}) : ``}`;
});

export { Layout as default };
//# sourceMappingURL=_layout.svelte-DJjCGr4B.js.map
