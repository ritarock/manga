import { GetServerSideProps } from "next";
import { BookData } from "../../interfaces";
import Viewer from "../../components/viewer"

interface Books {
  code: number
  data: BookData[]
}

export default function Manga({books}: {books: Books}) {
  return (
    <>
      <Viewer booksData={books.data} />
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async context => {
  const date = context.query.date
  const res = await fetch(`http://0.0.0.0:8080/backend/manga/release?date=${date}`)
  const books: Books = await res.json()

  return {
    props: {
      books,
    }
  }
}
