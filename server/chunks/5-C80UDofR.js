const API_URL$1 = "https://awcu2nks23.us-east-1.awsapprunner.com";
const load = async ({ params }) => {
  const response = await fetch(`${API_URL$1}/api/${params.event}/pools`);
  if (response.status == 404) {
    return {
      games: null
    };
  }
  const data = await response.json();
  return {
    games: data
  };
};

var _page_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const API_URL = "https://awcu2nks23.us-east-1.awsapprunner.com";
const actions = {
  update: async ({ cookies, request }) => {
    const data = await request.formData();
    const gameId = data.get("gameId");
    const team1Score = Number(data.get("team1Score"));
    const team2Score = Number(data.get("team2Score"));
    await fetch(`${API_URL}/api/update-pool`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ gameId, team1Score, team2Score })
    });
  },
  start: async ({ params }) => {
    await fetch(`${API_URL}/api/${params.event}/start-pools`, {
      method: "POST"
    });
  }
};

var _page_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  actions: actions
});

const index = 5;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-DXKTUpds.js')).default;
const universal_id = "src/routes/[event]/pools/+page.ts";
const server_id = "src/routes/[event]/pools/+page.server.ts";
const imports = ["_app/immutable/nodes/5.BUJHOCV3.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js","_app/immutable/chunks/TableHeadCell.DCsmDhEb.js","_app/immutable/chunks/bundle-mjs.C1WdQso_.js","_app/immutable/chunks/index.DGw2-XJO.js","_app/immutable/chunks/Modal.Yh2MhD28.js","_app/immutable/chunks/Wrapper.FS3LXdfC.js","_app/immutable/chunks/Frame.DXEIsCoV.js","_app/immutable/chunks/ToolbarButton.oLtMypMT.js","_app/immutable/chunks/Heading.BAZVKAx1.js","_app/immutable/chunks/EditOutline.Cu0SH5ig.js","_app/immutable/chunks/forms.CYrJzw3n.js","_app/immutable/chunks/entry.Cwgp6uEs.js","_app/immutable/chunks/admin.D1QBpGg2.js"];
const stylesheets = [];
const fonts = [];

export { component, fonts, imports, index, _page_server_ts as server, server_id, stylesheets, _page_ts as universal, universal_id };
//# sourceMappingURL=5-C80UDofR.js.map
