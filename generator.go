package protoc_gen_doc

//go:generate go build ./test/cmd/gen_fixtures
//go:generate protoc --plugin=protoc-gen-doc=./gen_fixtures --doc_out=. fixtures/Booking.proto fixtures/Vehicle.proto
//go:generate rm gen_fixtures
