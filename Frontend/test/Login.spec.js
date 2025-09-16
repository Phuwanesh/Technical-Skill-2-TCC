import { mount } from "@vue/test-utils";
import { describe, it, expect, vi } from "vitest";
import Login from "../src/components/Login.vue";

vi.mock("../src/api", () => ({
  default: {
    post: vi.fn((url, data) => {
      if (url === "/login") {
        if (data.username === "user" && data.password === "123456") {
          return Promise.resolve({ data: { token: "fake-jwt" } });
        } else {
          return Promise.reject({
            response: { data: { error: "invalid username or password" } },
          });
        }
      }
    }),
  },
}));

describe("Login.vue", () => {
  it("แสดง error ถ้า login ไม่สำเร็จ", async () => {
    const wrapper = mount(Login);

    await wrapper.find('input[type="text"]').setValue("wrong");
    await wrapper.find('input[type="password"]').setValue("wrong");
    await wrapper.find("form").trigger("submit.prevent");

    await new Promise(resolve => setTimeout(resolve, 0));

    expect(wrapper.text()).toContain("invalid username or password");
  });

  it("บันทึก token ถ้า login สำเร็จ", async () => {
    localStorage.clear();
    const wrapper = mount(Login);

    await wrapper.find('input[type="text"]').setValue("user");
    await wrapper.find('input[type="password"]').setValue("123456");
    await wrapper.find("form").trigger("submit.prevent");

    await new Promise(resolve => setTimeout(resolve, 0));

    expect(localStorage.getItem("token")).toBe("fake-jwt");
  });
});
