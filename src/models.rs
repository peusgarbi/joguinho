use crate::schema::saves;

#[derive(Insertable)]
#[table_name = "saves"]
pub struct NewSave<'a> {
    pub username: &'a str,
}

#[derive(Debug, Queryable, AsChangeset)]
pub struct Save {
    pub id: i32,
    pub username: String,
}