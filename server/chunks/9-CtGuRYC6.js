import { r as redirect, f as fail } from './index-DHSpIlkf.js';

const API_URL = "https://awcu2nks23.us-east-1.awsapprunner.com";
const load = async ({ params }) => {
  const response = await fetch(`${API_URL}/api/home`);
  const data = await response.json();
  return {
    events: data.events,
    year: data.year
  };
};

var _page_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const ADMIN_USERNAME = "username";
const ADMIN_PASSWORD = "password";
console.log("Admin username is set:", !!ADMIN_USERNAME);
console.log("Admin password is set:", !!ADMIN_PASSWORD);
const actions = {
  default: async ({ cookies, request }) => {
    console.log("Login attempt initiated");
    const data = await request.formData();
    const username = data.get("username");
    const password = data.get("password");
    console.log("Received username:", username);
    console.log("Password received:", !!password);
    if (username && password && username === ADMIN_USERNAME && password === ADMIN_PASSWORD) {
      cookies.set("session", "admin", { path: "/", httpOnly: false, sameSite: "strict", secure: false, maxAge: 60 * 60 * 24 * 7 });
      console.log("Credentials match, setting cookie");
      cookies.set("session", "admin", {
        path: "/",
        httpOnly: false,
        sameSite: "strict",
        maxAge: 60 * 60 * 24 * 7
      });
      console.log("Cookie set, redirecting");
      throw redirect(303, "/");
    }
    console.log("Login failed, returning error");
    return fail(400, { incorrect: true });
  }
};

var _page_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  actions: actions
});

const index = 9;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-DxFBY_Ac.js')).default;
const universal_id = "src/routes/home/+page.ts";
const server_id = "src/routes/home/+page.server.ts";
const imports = ["_app/immutable/nodes/9.lVAM-UG9.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js","_app/immutable/chunks/TableHeadCell.DCsmDhEb.js","_app/immutable/chunks/bundle-mjs.C1WdQso_.js","_app/immutable/chunks/Modal.Yh2MhD28.js","_app/immutable/chunks/Wrapper.FS3LXdfC.js","_app/immutable/chunks/Frame.DXEIsCoV.js","_app/immutable/chunks/ToolbarButton.oLtMypMT.js","_app/immutable/chunks/Tabs.Cg4KEYbt.js","_app/immutable/chunks/index.DGw2-XJO.js","_app/immutable/chunks/Heading.BAZVKAx1.js","_app/immutable/chunks/entry.Cwgp6uEs.js","_app/immutable/chunks/forms.CYrJzw3n.js"];
const stylesheets = [];
const fonts = [];

export { component, fonts, imports, index, _page_server_ts as server, server_id, stylesheets, _page_ts as universal, universal_id };
//# sourceMappingURL=9-CtGuRYC6.js.map
