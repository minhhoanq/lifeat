import type { Knex } from "knex";


export async function up(knex: Knex): Promise<void> {
    if (!(await knex.schema.hasTable(""))) {
        knex.schema.createTable("payments", (table) => {
            table.increments("id", { primaryKey: true });
            table.string("name", 256).nullable
        })
    }
}


export async function down(knex: Knex): Promise<void> {

}

