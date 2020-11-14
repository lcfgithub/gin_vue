<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col lg="6" offset-lg="3" md="8" offset-md="2">
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="请输入名称"
              ></b-form-input>
            </b-form-group>
            <b-form-valid-feedback :state="validateState('name')"
              >名称不能小于六个字符</b-form-valid-feedback
            >
            <!-- <b-form-text></b-form-text> -->
            <!-- <b-form-invalid-feedback :state="name_error">
                            不能小于六个字符
                        </b-form-invalid-feedback> -->
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telphone.$model"
                type="number"
                required
                placeholder="请输入名称"
              ></b-form-input>
            </b-form-group>
            <b-form-valid-feedback :state="validateState('telphone')"
              >手机号不合法</b-form-valid-feedback
            >

            <!-- <b-form-text
                            text-variant="danger"
                            v-if="isShowTelphone"
                        >
                            手机号必须11位
                        </b-form-text> -->

            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                required
                placeholder="请输入密码"
              ></b-form-input>
            </b-form-group>
            <b-form-valid-feedback :state="validateState('password')"
              >密码不能小于六个字符</b-form-valid-feedback
            >

            <b-form-group>
              <b-button @click="register" variant="outline-primary" block
                >注册</b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from "vuelidate/lib/validators";
import validator from "../../helper/validator";
import storageService from "../../service/storage";
import userService from "../../service/user";
// import { mapMutations } from "vuex";
import { createNamespacedHelpers, mapActions } from "vuex";
const { mapMutations } = createNamespacedHelpers("userModule");

export default {
  data() {
    return {
      user: {
        name: "",
        telphone: "",
        password: "",
      },
      isShowTelphone: false,
      name_error: null,
    };
  },
  validations: {
    user: {
      name: {
        required,
        minLength: minLength(6),
      },
      password: {
        required,
        minLength: minLength(6),
      },
      telphone: {
        telphone: validator.telphoneValidate,
      },
    },
  },
  methods: {
    // ...mapMutations("userModule", ["SET_TOKEN", "SET_USER_INFO"]),
    ...mapMutations(["SET_TOKEN", "SET_USER_INFO"]),
    ...mapActions("userModule", { userRegister: "register" }),
    validateState(field) {
      let { $dirty, $error } = this.$v.user[field];
      return $dirty ? $error : null;
    },
    register() {
      // this.isShowTelphone = this.user.telphone.length !== 11;
      // this.name_error = this.user.name ? this.user.name.length > 6 : null;

      //验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      } /*
      userService
        .register(this.user)
        .then((res) => {
          if (res.data.code != 200) {
            this.$bvToast.toast("用户已存在", {
              title: `温馨提示`,
              variant: "warning",
              solid: true,
            });
            return;
          }
          //保存token
          this.SET_TOKEN(res.data.data.token);
          // this.$store.commit("userModule/SET_TOKEN", res.data.data.token);
          // storageService._set(storageService.USER_TOKEN, res.data.data.token);
          return userService.info();
        })
        .then((resp) => {
          // storageService._set(
          //   storageService.USER_INFO,
          //   JSON.stringify(resp.data.user)
          // );
          // this.$store.commit("userModule/SET_USER_INFO", resp.data.user);
          this.SET_USER_INFO(resp.data.user);

          //跳转主页
          this.$router.replace({ name: "Home" });
        })
        .catch((err) => {
          this.$bvToast.toast(err, {
            title: `网络异常`,
            variant: "danger",
            solid: true,
          });
          console.log(err);
        });
        */
      /*
      userService
        .register(this.user)
        .then((res) => {
          //保存token
          console.log(res.data);
          if (res.data.code != 200) {
            this.$bvToast.toast("用户已存在", {
              title: `温馨提示`,
              variant: "warning",
              solid: true,
            });
            return;
          }
          // localStorage.setItem("_token", res.data.data.token);
          storageService._set(storageService.USER_TOKEN, res.data.data.token);
          userService.info().then((resp) => {
            storageService._set(
              storageService.USER_INFO,
              JSON.stringify(resp.data.user)
            );
            //跳转主页
            this.$router.replace({ name: "Home" });
            window.location.reload();
          });
        })
        .catch((err) => {
          this.$bvToast.toast(err, {
            title: `网络异常`,
            variant: "danger",
            solid: true,
          });
          console.log(err);
        });
        */
      this.userRegister(this.user) //this.$store.dispatch("userModule/register", this.user)
        .then((res) => {
          console.log("userModule/register", res.data.code, res);
          if (res.data.user && res.status == 200) {
            this.$router.replace({ name: "Home" });
          }
          if (res.data.code != 200) {
            this.$bvToast.toast(res.data.msg, {
              title: `温馨提示`,
              variant: "warning",
              solid: true,
            });
          }
        })
        .catch((err) => {
          this.$bvToast.toast(JSON.stringify(err), {
            title: `网络异常`,
            variant: "danger",
            solid: true,
          });
          console.log(err);
        });
    },
  },
};
</script>
