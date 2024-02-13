import { Module } from '@nestjs/common';
import { ContainersService } from './containers.service';
import { ContainersController } from './containers.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Container } from './entities/container.entity';
import { FeatureFlagsModule } from 'src/feature-flags/feature-flags.module';

@Module({
  imports: [TypeOrmModule.forFeature([Container]), FeatureFlagsModule],
  controllers: [ContainersController],
  providers: [ContainersService],
  exports: [ContainersService],
})
export class ContainersModule {}
