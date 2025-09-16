import { mount } from "@vue/test-utils";
import { describe, it, expect, vi } from "vitest";
import Register from "../src/components/Register.vue";
import api from "../src/api";

vi.mock("../src/api"); // mock axios instance

describe("Register.vue", () => {
  it("สมัครสมาชิกสำเร็จ", async () => {
    api.post.mockResolvedValueOnce({ data: { message: "registered" } });

    const wrapper = mount(Register);

    await wrapper.find('input[type="text"]').setValue("newuser");
    await wrapper.find('input[type="password"]').setValue("123456");
    await wrapper.findAll('input[type="password"]')[1].setValue("123456");

    await wrapper.find("form").trigger("submit.prevent");

    // ในโค้ดคุณใช้ alert → mock ไว้ก่อน
    expect(api.post).toHaveBeenCalledWith("/register", {
      username: "newuser",
      password: "123456",
      confirmPassword: "123456",
    });
  });

  it("สมัครไม่สำเร็จ (username ซ้ำ)", async () => {
    api.post.mockRejectedValueOnce({
      response: { data: { error: "username already exists" } },
    });

    const wrapper = mount(Register);

    await wrapper.find('input[type="text"]').setValue("user1");
    await wrapper.find('input[type="password"]').setValue("123456");
    await wrapper.findAll('input[type="password"]')[1].setValue("123456");

    await wrapper.find("form").trigger("submit.prevent");

    expect(wrapper.html()).toContain("username already exists");
  });

  it("สมัครไม่สำเร็จ (password ไม่ตรง)", async () => {
    api.post.mockRejectedValueOnce({
      response: { data: { error: "password and confirmPassword do not match" } },
    });

    const wrapper = mount(Register);

    await wrapper.find('input[type="text"]').setValue("user2");
    await wrapper.find('input[type="password"]').setValue("123456");
    await wrapper.findAll('input[type="password"]')[1].setValue("654321");

    await wrapper.find("form").trigger("submit.prevent");

    expect(wrapper.html()).toContain("password and confirmPassword do not match");
  });
});
