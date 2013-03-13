#import <Foundation/NSObject.h>

@interface Hello
{
@public int publicInt;
}
+ classMethod;
+ classMethod2;
- objMethod1;
- objMethod2;
@end

@interface Hello2
{
@public int publicHello2Int;
}
+ hello2ClassMethod;
- hello2ClassMethod;
@end


@interface World
{
@public
    Hello2* worldVar;
}
- (Hello*) world;
- (void) setWorld:(Hello*) world;
@end

@interface World2 : NSObject
- (World*) world2;
- (void) setWorld2:(World*) world2;
@end


@interface World3: NSObject
{
@private
    int privateVariable;
@protected
    int protectedVariable;
@public
    int publicVariable;

}
@property float value;
- (void) something;

@end

@implementation World3
- (void) something
{
    World3* w;
}
@end

@interface World4 : World3
- (void) myworld;
@end

@implementation World4
- (void) myworld
{

}
@end

@interface World4 (MyCategory)
- (void) MyCategoryMethod;
@end

@protocol Serializeable
-(void) serialize;
-(void) deserialize;
@end

@interface World5 : World4<Serializeable>
@end

int main(int argc, char const *argv[])
{
	return 0;
}
