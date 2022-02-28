<template>
  <div>
    <h1>Step 6</h1>
    <!-- Table to display data -->
    <table class="table table-bordered" id="usersTable">
      <tbody>
        <tr>
          <th>Reference No</th>
          <th>Name</th>
          <th>Control</th>
        </tr>
        <tr v-for="(item, index) in users" :key="item">
          <td class="refno">{{ index }}</td>
          <td>
            <button
              type="button"
              class="btn btn-info btn-lg"
              data-toggle="modal"
              data-target="#user"
              v-on:click="getreq({ index })"
            >
              {{ item }}
            </button>
          </td>
          <td colspan="2">
            <button
              type="button"
              class="btn btn-info btn-lg"
              data-toggle="modal"
              data-target="#edit"
            >
              Edit
            </button>
            <button
              type="button"
              class="btn btn-info btn-lg"
              data-toggle="modal"
              data-target="#delete"
            >
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <!-- Modal Window  for users  -->
    <div id="user" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <!-- Modal content-->
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              &times;
            </button>
            <h4 class="modal-title">User Data</h4>
          </div>
          <div class="modal-body">
            <!-- Table to display data -->
            <table class="table">
              <tbody id="table_user">
                <tr>
                  <td>
                    Name : {{ users_value.user_name }}
                  </td>
                  <td>
                   Team : {{ users_value.team_name }}
                  </td>
                  <td>
                    Location : {{ users_value.location_name }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="modal-footer">
            <button
              type="button"
              class="btn btn-info btn-lg"
              data-toggle="modal"
              data-target="#edit"
            >
              Edit
            </button>
            <button
              type="button"
              class="btn btn-info btn-lg"
              data-toggle="modal"
              data-target="#delete"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "main_html",
  data() {
    return {
      users: [],
      users_value: [],
    };
  },
  methods: {
    getreqmul() {
      const url = "http://localhost:8080";
      fetch(url)
        .then((res) => res.json())
        .then((data) => {
          this.users = data.data;
        })
        .catch((err) => console.error(err));
    },
    getreq(index) {
      let i = Object.values(index);
      let url = "http://localhost:8080/users?ref=" + i[0];
      fetch(url)
        .then((res) => res.json())
        .then((data) => {
          this.users_value = data.data;
        })
        .catch((err) => console.error(err));
    },
  },
  mounted() {
    this.getreqmul();
  },
};
</script>
