import { i as isAdmin } from './admin-DrRi-8BF.js';
import './index2-CpnJNRxb.js';
import './utils-Cb-NyS5y.js';

const load$1 = async ({ data }) => {
  isAdmin.set(data.isAdmin);
};

var _layout_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load$1
});

const load = async ({ cookies }) => {
  const isAdmin = cookies.get("session") === "admin";
  return { isAdmin };
};

var _layout_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const index = 0;
let component_cache;
const component = async () => component_cache ??= (await import('./_layout.svelte-Iok0mYdi.js')).default;
const universal_id = "src/routes/+layout.ts";
const server_id = "src/routes/+layout.server.ts";
const imports = ["_app/immutable/nodes/0.BilmvqqF.js","_app/immutable/chunks/admin.D1QBpGg2.js","_app/immutable/chunks/index.DGw2-XJO.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js","_app/immutable/chunks/bundle-mjs.C1WdQso_.js","_app/immutable/chunks/Frame.DXEIsCoV.js","_app/immutable/chunks/ToolbarButton.oLtMypMT.js","_app/immutable/chunks/Heading.BAZVKAx1.js"];
const stylesheets = ["_app/immutable/assets/0.DoYQhQLu.css"];
const fonts = [];

export { component, fonts, imports, index, _layout_server_ts as server, server_id, stylesheets, _layout_ts as universal, universal_id };
//# sourceMappingURL=0-CWKxTb9C.js.map
