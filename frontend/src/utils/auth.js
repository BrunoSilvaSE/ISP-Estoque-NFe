import { replaceByRole } from "/src/utils/replaceByRole";

(async function () {
  const token = localStorage.getItem("token");
  if (!token) {
    await replaceByRole(token);
  }else {
    window.location.replace("/")
  }
})();
