#include <stdint.h>
#include <stdio.h>
#include "cppgo.h"

extern "C" void TestFoo(uintptr_t foo) {
  go::demo::Foo f(reinterpret_cast<const void*>(foo));
  for (auto b : f.slice()) {
    if (!b) {
      printf("null pointer\n");
    } else {
      printf("%lld %f %s\n", b->int_val(), b->float_val(),
             b->string_val().as_string().c_str());
    }
  }
}
